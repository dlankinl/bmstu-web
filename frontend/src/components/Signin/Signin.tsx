import React, { useState } from 'react';
import Form from '../Form/Form';
import { signinFields } from '../Form/formConfigs';
import ErrorPopup from '../ErrorPopup/ErrorPopUp';

const BASE_URL = 'http://localhost:8081/api/v2/login';



const SigninComponent: React.FC<{ onSuccess: () => void }> = ({ onSuccess }) => {
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [showErrorPopup, setShowErrorPopup] = useState(false);

  const handleCloseError = () => {
    setShowErrorPopup(false);
    setErrorMessage(null);
  };

  const handleSubmit = async (data: any) => {
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

    try {
      const response = await fetch(BASE_URL, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          login: data.username,
          password: data.password,
        }),
      });

      if (!response.ok) {
        throw new Error('Ошибка входа. Проверьте ваши учетные данные.');
      }

      const responseData = await response.json();
      const token = responseData.token;

      localStorage.setItem('authToken', token);

      console.log('Login successful:', responseData);
      // alert('Вы успешно вошли в систему!');
      onSuccess();
    } catch (error) {
      setErrorMessage(error.message);
      setShowErrorPopup(true);
    }
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
