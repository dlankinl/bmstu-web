import React from 'react';
import Form from '../Form/Form';
import { companyFields } from '../Form/formConfigs';

const CreateCompanyComponent = ({ }) => {
  const handleSubmit = (data: any) => {
    if (data.name === '') {
      alert("Введите название!");
      return;
    }

    if (data.city === '') {
      alert("Введите название города!");
      return;
    }

    if (data.activityFieldID === '') {
      alert("Выберите сферу деятельности!");
      return;
    }

    console.log(data); 
    alert(JSON.stringify(data));
  };

  return (
    <Form title="Создать компанию" buttonText='' fields={companyFields} onSubmit={handleSubmit} />
  );
};

export default CreateCompanyComponent;
