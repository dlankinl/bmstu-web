import React, { useState } from 'react';
import Form from '../Form/Form';
import { financialReportFields } from '../Form/formConfigs';
import ErrorPopup from '../ErrorPopup/ErrorPopUp';

function isNumeric(value: any) {
  return /^\d+(\.\d+)?$/.test(value);
}

const CreateFinancialReportComponent = ({ }) => {
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [showErrorPopup, setShowErrorPopup] = useState(false);

  const handleCloseError = () => {
    setShowErrorPopup(false);
    setErrorMessage(null);
  };

  const handleSubmit = (data: any) => {
    if (!isNumeric(data.startYear) || data.startYear < 0) {
      setErrorMessage("Введите год начала периода (целое положительное число)!");
      setShowErrorPopup(true);
      return;
    }

    if (!isNumeric(data.startQuarter) || data.startQuarter < 1 || data.startQuarter > 4) {
      setErrorMessage("Введите квартал начала периода (целое положительное число от 1 до 4 включительно)!");
      setShowErrorPopup(true);
      return;
    }

    if (!isNumeric(data.endYear) || data.endYear < 0) {
      setErrorMessage("Введите год конца периода (целое положительное число)!");
      setShowErrorPopup(true);
      return;
    }

    if (data.endYear < data.startYear) {
      setErrorMessage("Год конца периода должен быть больше либо равен году начала!");
      setShowErrorPopup(true);
      return;
    }

    if (!isNumeric(data.endQuarter) || data.endQuarter < 1 || data.endQuarter > 4) {
      setErrorMessage("Введите квартал конца периода (целое положительное число от 1 до 4 включительно)!");
      setShowErrorPopup(true);
      return;
    }

    if (!isNumeric(data.revenue) || data.revenue < 0) {
      setErrorMessage("Введите сумму выручки (положительное число)!");
      setShowErrorPopup(true);
      return;
    }

    if (!isNumeric(data.costs) || data.costs < 0) {
      setErrorMessage("Введите сумму расходов (положительное число)!");
      setShowErrorPopup(true);
      return
    }

    console.log(data); 
    alert(JSON.stringify(data));
  };

  return (
    <>
      <Form title="Добавить финансовый отчет" buttonText='' fields={financialReportFields} onSubmit={handleSubmit} />
      {showErrorPopup && errorMessage && (
        <ErrorPopup message={errorMessage} onClose={handleCloseError} />
      )}
    </>
  );
};

export default CreateFinancialReportComponent;
