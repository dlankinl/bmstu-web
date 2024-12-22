import React, { useState } from 'react';
import Form from '../Form/Form';
import { companyFields } from '../Form/formConfigs';
import ErrorPopup from '../ErrorPopup/ErrorPopUp';

interface CreateCompanyComponentProps {
  isEditing?: boolean;
}

const EditCompanyComponent = ({   }) => {
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [showErrorPopup, setShowErrorPopup] = useState(false);

  const handleCloseError = () => {
    setShowErrorPopup(false);
    setErrorMessage(null);
  };

  const handleSubmit = (data: any) => {
    if (data.name === '') {
      setErrorMessage("Введите название!");
      setShowErrorPopup(true);
      return;
    }

    if (data.city === '') {
      setErrorMessage("Введите название города!");
      setShowErrorPopup(true);
      return;
    }

    if (data.activityFieldID === '') {
      setErrorMessage("Выберите сферу деятельности!");
      setShowErrorPopup(true);
      return;
    }

    console.log(data); 
    alert(JSON.stringify(data));
  };

  return (
    <>  
      <Form title="Редактировать компанию" buttonText='Сохранить' fields={companyFields} onSubmit={handleSubmit} />
      {showErrorPopup && errorMessage && (
        <ErrorPopup message={errorMessage} onClose={handleCloseError} />
      )}
    </>
  );
};

export default EditCompanyComponent;
