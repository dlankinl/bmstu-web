import React from 'react';
import Form from '../Form/Form';
import { activityFieldFields } from '../Form/formConfigs';

const CreateActivityFieldComponent = ({ }) => {
  const handleSubmit = (data: any) => {
    console.log(data); 
    alert(JSON.stringify(data));
  };

  return (
    <Form title="Создать сферу деятельности" buttonText='' fields={activityFieldFields} onSubmit={handleSubmit} />
  );
};

export default CreateActivityFieldComponent;
