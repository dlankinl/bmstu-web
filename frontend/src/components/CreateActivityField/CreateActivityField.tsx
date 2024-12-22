import React from 'react';
import Form from '../Form/Form';
import { activityFieldFields } from '../Form/formConfigs';

function isNumeric(value: any) {
  return /^\d+(\.\d+)?$/.test(value);
}

const CreateActivityFieldComponent = ({ }) => {
  const handleSubmit = (data: any) => {
    if (data.name === '') {
      alert("Введите название!");
      return;
    }

    if (data.description === '') {
      alert("Введите описание!");
      return;
    }

    if (data.cost === '') {
      alert("Введите вес сферы деятельности!");
      return;
    }

    if (!isNumeric(data.cost)) {
      alert("Введите положительное число!");
      return
    }

    console.log(data); 
    alert(JSON.stringify(data));
  };

  return (
    <Form title="Создать сферу деятельности" buttonText='' fields={activityFieldFields} onSubmit={handleSubmit} />
  );
};

export default CreateActivityFieldComponent;
