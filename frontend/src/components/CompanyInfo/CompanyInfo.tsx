import React from 'react';
import { useParams } from 'react-router-dom';
import { useState, useEffect } from 'react';
import './CompanyInfo.css';
import GreenButton from '../Buttons/GreenButton';
import FormContainer from '../Form/FormContainer';
import { CompanyInfo } from './types';
import Modal from '../Modal/Modal';
import FinancialReportComponent from '../FinancialReport/FinancialReport';
import CreateFinancialReportComponent from '../CreateFinancialReport/CreateFinancialReport';
import CreateCompanyComponent from '../CreateCompany/CreateCompany';
import FinancialReportForm from '../FinancialReportForm/FinancialReportForm';
import { ActivityField } from '../ActivityFields/types';
import { Entrepreneur } from '../Entrepreneurs/types';

const BASE_URL = 'http://localhost:8081/api/v2/companies';
const BASE_URL_ACT_FIELDS = 'http://localhost:8081/api/v2/activity_fields';
const BASE_URL_ENTS = 'http://localhost:8081/api/v2/entrepreneurs';

const CompanyInfoComponent = ({ }) => {
  const { id } = useParams();
  const [values, setValues] = useState<CompanyInfo>();
  const [actFieldValue, setActField] = useState<ActivityField>();
  const [entValue, setEntValue] = useState<Entrepreneur>();
  const [loading, setLoading] = useState(true);
  
  const [isModalActive, setModalActive] = useState(false);
  const [modalContent, setModalContent] = useState<React.ReactNode>(null);

  const handleModalOpen = (content: React.ReactNode) => {
    setModalContent(content);
    setModalActive(true);
  };

  const handleModalClose = () => {
    setModalActive(false);
    setModalContent(null);
  };

  var isOwner = false;

  useEffect(() => {
    const fetchCompanyInfo = async () => {
      try {
        const response = await fetch(`${BASE_URL}/${id}`);
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data: CompanyInfo = await response.json();
        setValues(data.company);
        
        await Promise.all([
          fetchActivityFieldName(data.company.activityFieldId),
          fetcEntrepreneurName(data.company.ownerId),
        ]);
      } catch (error) {
        console.error("Error fetching entrepreneur data:", error);
      } finally {
        setLoading(false);
      }
    };
  
    const fetchActivityFieldName = async (activityFieldId: string) => { 
      try {
        console.log("Fetching activity field with ID:", activityFieldId);
        const response = await fetch(`${BASE_URL_ACT_FIELDS}/${activityFieldId}`);
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const actFieldData: ActivityField = await response.json();
        setActField(actFieldData.activityField);
      } catch (error) {
        console.error("Error fetching activity field data: ", error);
      } 
    };

    const fetcEntrepreneurName = async (ownerId: string) => { 
      try {
        console.log("Fetching entrepreneur with ID:", ownerId);
        const response = await fetch(`${BASE_URL_ENTS}/${ownerId}`);
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const entValue: Entrepreneur = await response.json();
        setEntValue(entValue.entrepreneur);
      } catch (error) {
        console.error("Error fetching entrepreneur data: ", error);
      } 
    };
  
    fetchCompanyInfo();
  }, [id]);

  const parts = localStorage.getItem("authToken")?.split('.');
  if (parts !== undefined) {
    const decodedPayload = atob(parts[1]);

    const parsedPayload = JSON.parse(decodedPayload);
    isOwner = parsedPayload.sub === values?.ownerId;
  }
  

  const handleFinancialReportSubmit = (data: { startYear: number; endYear: number; startQuarter: number; endQuarter: number }) => {
    const reportResults = (
      <FinancialReportComponent params={data} />
    );

    handleModalOpen(reportResults);
  };

  return (
    <FormContainer>
      <div className="my-component">
        <div className="row">
          <span>Название</span>
          <span>{values?.name}</span>
        </div>
        <div className="row">
          <span>Город</span>
          <span>{values?.city}</span>
        </div>
        <div className="row">
          <span>Сфера деятельности</span>
          <span>{actFieldValue?.name}</span>
        </div>
        <div className="row">
          <span>Владелец</span>
          <span>{entValue?.fullName}</span>
        </div>

        <div className="button-row">
          {isOwner && (
            <>
              <div className="button-wrapper">
                <GreenButton text={"Редактировать"} icon={""} onClick={() => { handleModalOpen(<CreateCompanyComponent isEditing={true} />) }} />
              </div>
              <div className="button-wrapper">
                <GreenButton text={"Добавить фин. отчет"} icon={""} onClick={() => { handleModalOpen(<CreateFinancialReportComponent />) }} />
              </div>
            </> 
          )}
        </div>

        <div className="button-wrapper-center">
          <div className="button-wrapper">
            <GreenButton text={"Финансовый отчет"} icon={""} onClick={() => { handleModalOpen(<FinancialReportForm onSubmit={handleFinancialReportSubmit} />) }} />
          </div>
        </div>

      </div>

      {/* Modal for displaying content */}
      {isModalActive && (
        <Modal onClose={handleModalClose}>
          {modalContent}
        </Modal>
      )}
    </FormContainer>
  );
};

export default CompanyInfoComponent;
