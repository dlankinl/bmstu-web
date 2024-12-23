import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import './EntrepreneurInfo.css';
import GreenButton from '../Buttons/GreenButton';
import FormContainer from '../Form/FormContainer';
import { EntrepreneurInfo } from './types';
import Modal from '../Modal/Modal';
import FinancialReportComponent from '../FinancialReport/FinancialReport';
import EditEntrepreneurComponent from '../EditEntrepreneur/EditEntrepreneur';
import FinancialReportForm from '../FinancialReportForm/FinancialReportForm';
import ContactsList from '../ContactsList/ContactsList';

const BASE_URL = 'http://localhost:8081/api/v2/entrepreneurs';

const EntrepreneurInfoComponent: React.FC = () => {
  const { id } = useParams();
  const [values, setValues] = useState<EntrepreneurInfo>();
  const [isModalActive, setModalActive] = useState(false);
  const [modalContent, setModalContent] = useState<React.ReactNode>(null);
  const [loading, setLoading] = useState(true); // Loading state

  const handleModalOpen = (content: React.ReactNode) => {
    setModalContent(content);
    setModalActive(true);
  };

  const handleModalClose = () => {
    setModalActive(false);
    setModalContent(null);
  };

  useEffect(() => {
    const fetchEntrepreneurInfo = async () => {
      try {
        const response = await fetch(`${BASE_URL}/${id}`);
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data: EntrepreneurInfo = await response.json();
        setValues(data.entrepreneur); // Assuming data.entrepreneur contains the relevant info
      } catch (error) {
        console.error("Error fetching entrepreneur data:", error);
      } finally {
        setLoading(false); // Set loading to false after fetching
      }
    };

    fetchEntrepreneurInfo();
  }, [id]); 

  const navigate = useNavigate(); // Initialize the navigate function

  const navigateToCompanies = () => {
    navigate(`/entrepreneurs/${values?.id}/companies`);
  };
  const isAuthenticated = !!localStorage.getItem('authToken');

  const handleFinancialReportSubmit = (data: { startYear: number; endYear: number; startQuarter: number; endQuarter: number }) => {
    const reportResults = (
      <FinancialReportComponent params={data} />
    );

    handleModalOpen(reportResults);
  };

  if (loading) {
    return <div>Загрузка...</div>;
  }

  return (
    <FormContainer>
      <div className="my-component">
        <div className="row">
          <span>ФИО</span>
          <span>{values?.fullName}</span>
        </div>
        <div className="row">
          <span>Город</span>
          <span>{values?.city}</span>
        </div>
        <div className="row">
          <span>Дата рождения</span>
          <span>{values?.birthday}</span>
        </div>
        <div className="row">
          <span>Пол</span>
            <span>
              {values?.gender === 'm' ? 'мужской' : 
              values?.gender === 'w' ? 'женский' : 
              values?.gender ? values.gender : 'неизвестно'}
            </span>
        </div>

        <div className="button-row">
          <div className="button-wrapper">
            <GreenButton text={"Редактировать"} icon={""} onClick={() => { handleModalOpen(<EditEntrepreneurComponent />) }} />
          </div>
          <div className="button-wrapper">
            <GreenButton text={"Финансовый отчет"} icon={""} onClick={() => { handleModalOpen(<FinancialReportForm onSubmit={handleFinancialReportSubmit} />) }} />
          </div>
        </div>

        {!isAuthenticated ? (
          <div className="button-wrapper-center">
            <div className="button-wrapper">
                <GreenButton text={"Компании"} icon={""} onClick={navigateToCompanies}/>
            </div>
          </div>
        ) : (
          <div className="button-row">
            <div className="button-wrapper">
                <GreenButton text={"Компании"} icon={""} onClick={navigateToCompanies}/>
            </div>
            <div className="button-wrapper">
              <GreenButton text={"Контакты"} icon={""} onClick={() => { handleModalOpen(<ContactsList />) }} />
            </div>
          </div>
        )}
      </div>

      {isModalActive && (
        <Modal onClose={handleModalClose}>
          {modalContent}
        </Modal>
      )}
    </FormContainer>
  );
};

export default EntrepreneurInfoComponent;
