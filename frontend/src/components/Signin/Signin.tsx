import React, { useState } from 'react';
import Form from '../Form/Form';
import { signinFields } from '../Form/formConfigs';
import ErrorPopup from '../ErrorPopup/ErrorPopUp';

const SigninComponent = ({ }) => {
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [showErrorPopup, setShowErrorPopup] = useState(false);

  const handleCloseError = () => {
    setShowErrorPopup(false);
    setErrorMessage(null);
  };

  const handleSubmit = (data: any) => {
    if (data.username === '') {
      setErrorMessage("Введите имя пользователя!");
      setShowErrorPopup(true);
      return;
    }

    if (data.password === '') {
      setErrorMessage("Введите пароль!");
      setShowErrorPopup(true);
      return;
    }

    if (data.password.length < 8) {
      setErrorMessage("Пароль должен состоять не менее, чем из 8 символов!");
      setShowErrorPopup(true);
      return;
    }

    console.log(data); 
    alert(JSON.stringify(data));
  };

  return (
    <>
      <Form title="Аутентификация" fields={signinFields} buttonText='Войти' onSubmit={handleSubmit} />
      {showErrorPopup && errorMessage && (
        <ErrorPopup message={errorMessage} onClose={handleCloseError} />
      )}
    </>
  );
};

export default SigninComponent;
