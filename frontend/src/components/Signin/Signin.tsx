import React from 'react';
import Form from '../Form/Form';
import { signinFields } from '../Form/formConfigs';

const SigninComponent = ({ }) => {
  const handleSubmit = (data: any) => {
    if (data.username === '') {
      alert("Укажите имя пользователя!");
      return;
    }

    if (data.password === '') {
      alert("Укажите пароль!");
      return;
    }

    if (data.password.length < 8) {
      alert("Пароль должен состоять не менее, чем из 8 символов!");
      return;
    }

    console.log(data); 
    alert(JSON.stringify(data));
  };

  return (
    <Form title="Аутентификация" fields={signinFields} buttonText='Войти' onSubmit={handleSubmit} />
  );
};

export default SigninComponent;
