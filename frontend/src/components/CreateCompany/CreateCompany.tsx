import React, { useState } from 'react';
import Form from '../Form/Form';
import { companyFields } from '../Form/formConfigs';
import ErrorPopup from '../ErrorPopup/ErrorPopUp';

interface CreateCompanyComponentProps {
  isEditing?: boolean;
}

const CreateCompanyComponent: React.FC<CreateCompanyComponentProps> = ({ isEditing }) => {
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
      setErrorMessage("Введите сферу деятельности!");
      setShowErrorPopup(true);
      return;
    }

    console.log(data); 
    alert(JSON.stringify(data));
  };

  const buttonText: string = isEditing ? "Сохранить" : "Добавить";
  const title: string = isEditing ? "Редактировать компанию" : "Создать компанию";

  return (
    <>
      <Form title={title} buttonText={buttonText} fields={companyFields} onSubmit={handleSubmit} />
      {showErrorPopup && errorMessage && (
        <ErrorPopup message={errorMessage} onClose={handleCloseError} />
      )}
    </>
  );
};

export default CreateCompanyComponent;
