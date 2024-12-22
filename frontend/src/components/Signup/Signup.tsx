import React from 'react';
import Form from '../Form/Form';
import { signupFields } from '../Form/formConfigs';

const SignupComponent = ({ }) => {
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

    if (data.verifyPassword === '') {
      alert("Подвердите пароль!");
      return;
    }

    if (data.password !== data.verifyPassword) {
      alert("Пароли должны совпадать!");
      return;
    }
    console.log(data); 
    alert(JSON.stringify(data));
  };

  return (
    <Form title="Регистрация" fields={signupFields} buttonText='Зарегистрироваться' onSubmit={handleSubmit} />
  );
};

export default SignupComponent;
