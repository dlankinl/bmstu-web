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

const CompanyInfoComponent = ({ }) => {
  const { id } = useParams();
  const [values, setValues] = useState<CompanyInfo>();

  const [isModalActive, setModalActive] = useState(false);
  const [modalContent, setModalContent] = useState<React.ReactNode>(null); // State to track modal content

  const handleModalOpen = (content: React.ReactNode) => {
    setModalContent(content); // Set the content based on button clicked
    setModalActive(true);
  };

  const handleModalClose = () => {
    setModalActive(false);
    setModalContent(null); // Reset content when closing
};

  useEffect(() => {
    const fetchCompanyInfo = () => {
      const staticCompanyInfo: CompanyInfo = { ID: "1", Name: "Test1", OwnerID: "1", City: "Moscow", ActivityFieldID: "1" };
      setValues(staticCompanyInfo);
    };

    // const fetchData = async () => {
    //   try {
    //     const response = await fetch(`https://api.example.com/companies/${id}`);
    //     const data = await response.json();
    //     setValues(data);
    //   } catch (error) {
    //     console.error("Error fetching company data:", error);
    //   }
    // };

    fetchCompanyInfo();
  }, [id]); 

  return (
    <FormContainer>
      <div className="my-component">
        <div className="row">
          <span>Название</span>
          <span>{values?.Name}</span>
        </div>
        <div className="row">
          <span>Город</span>
          <span>{values?.City}</span>
        </div>
        <div className="row">
          <span>Сфера деятельности</span>
          <span>{values?.ActivityFieldID}</span>
        </div>
        <div className="row">
          <span>Владелец</span>
          <span>{values?.OwnerID}</span>
        </div>

        <div className="button-row">
          <div className="button-wrapper">
            <GreenButton text={"Редактировать"} icon={""}/>
          </div>
          <div className="button-wrapper">
            {/* <GreenButton text={"Финансовый отчет"} icon={""} onClick={handleModalOpen}/> */}
            <GreenButton text={"Финансовый отчет"} icon={""} onClick={() => { handleModalOpen(<FinancialReportComponent/>) }}/>
          </div>
        </div>
        <div className="button-wrapper-center">
          <div className="button-wrapper">
            {/* <GreenButton text={"Добавить фин. отчет"} icon={""}/> */}
            <GreenButton text={"Добавить фин. отчет"} icon={""} onClick={() => { handleModalOpen(<CreateFinancialReportComponent/>) }}/>
          </div>
        </div>
      </div>
      <div>
        {isModalActive && (
          <Modal onClose={handleModalClose}>
            {modalContent}
          </Modal>
        )}
      </div>
    </FormContainer>
  );
};

export default CompanyInfoComponent;
