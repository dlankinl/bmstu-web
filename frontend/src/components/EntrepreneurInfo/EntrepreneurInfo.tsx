import React from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import './EntrepreneurInfo.css';
import GreenButton from '../Buttons/GreenButton';
import FormContainer from '../Form/FormContainer';
import { EntrepreneurInfo } from './types';
import Modal from '../Modal/Modal';
import FinancialReportComponent from '../FinancialReport/FinancialReport';
import EditEntrepreneurComponent from '../EditEntrepreneur/EditEntrepreneur';
import FinancialReportForm from '../FinancialReportForm/FinancialReportForm';

const EntrepreneurInfoComponent = ({ }) => {
  const { id } = useParams();
  const [values, setValues] = useState<EntrepreneurInfo>();

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

  useEffect(() => {
    const fetchEntrepreneurInfo = () => {
      const staticEntrepreneurInfo: EntrepreneurInfo = { ID: "1", Name: "Test1", Birthday: "22 декабря 2000 г.", City: "Moscow", Gender: "мужской" };
      setValues(staticEntrepreneurInfo);
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

    fetchEntrepreneurInfo();
  }, [id]); 

  const navigate = useNavigate(); // Initialize the navigate function

  const handleButtonClick = () => {
    navigate(`/entrepreneurs/${values?.ID}/companies`);
  };

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
          <span>ФИО</span>
          <span>{values?.Name}</span>
        </div>
        <div className="row">
          <span>Город</span>
          <span>{values?.City}</span>
        </div>
        <div className="row">
          <span>Дата рождения</span>
          <span>{values?.Birthday}</span>
        </div>
        <div className="row">
          <span>Пол</span>
          <span>{values?.Gender}</span>
        </div>

        <div className="button-row">
          <div className="button-wrapper">
            <GreenButton text={"Редактировать"} icon={""} onClick={() => { handleModalOpen(<EditEntrepreneurComponent />) }}
            />
          </div>
          <div className="button-wrapper">
            <GreenButton text={"Финансовый отчет"} icon={""} onClick={() => { handleModalOpen(<FinancialReportForm onSubmit={handleFinancialReportSubmit} />) }} />
          </div>
        </div>

        <div className="button-row">
         <div className="button-wrapper">
            <GreenButton text={"Компании"} icon={""} onClick={handleButtonClick}/>
          </div>
          <div className="button-wrapper">
            <GreenButton text={"Контакты"} icon={""} onClick={handleButtonClick}/>
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

export default EntrepreneurInfoComponent;
