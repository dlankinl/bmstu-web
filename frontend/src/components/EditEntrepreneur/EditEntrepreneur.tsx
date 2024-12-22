import React, { useState } from 'react';
import Form from '../Form/Form';
import { entrepreneurFields } from '../Form/formConfigs';
import ErrorPopup from '../ErrorPopup/ErrorPopUp';

const EditEntrepreneurComponent = ({ }) => {
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [showErrorPopup, setShowErrorPopup] = useState(false);

  const handleCloseError = () => {
    setShowErrorPopup(false);
    setErrorMessage(null);
  };

  const handleSubmit = (data: any) => {
    console.log(data)
    if (data.name === '') {
      setErrorMessage("Введите ФИО!");
      setShowErrorPopup(true);
      return;
    }

    if (data.gender === '') {
      setErrorMessage("Введите пол!");
      setShowErrorPopup(true);
      return;
    }

    if (data.birthday === '') {
      setErrorMessage("Введите дату рождения!");
      setShowErrorPopup(true);
      return;
    }

    if (data.city === '') {
      setErrorMessage("Введите название города!");
      setShowErrorPopup(true);
      return;
    }

    console.log(data); 
    alert(JSON.stringify(data));
  };

  return (
    <>
      <Form title="Редактирование профиля" buttonText="Сохранить" fields={entrepreneurFields} onSubmit={handleSubmit} />
      {showErrorPopup && errorMessage && (
        <ErrorPopup message={errorMessage} onClose={handleCloseError} />
      )}
    </>
  );
};

export default EditEntrepreneurComponent;
