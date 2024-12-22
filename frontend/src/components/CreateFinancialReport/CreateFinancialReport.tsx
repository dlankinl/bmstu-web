import React from 'react';
import Form from '../Form/Form';
import { financialReportFields } from '../Form/formConfigs';

function isNumeric(value: any) {
  return /^\d+(\.\d+)?$/.test(value);
}

const CreateFinancialReportComponent = ({ }) => {
  const handleSubmit = (data: any) => {
    if (!isNumeric(data.startYear) || data.startYear < 0) {
      alert("Введите год начала периода (целое положительное число)!");
      return;
    }

    if (!isNumeric(data.startQuarter) || data.startQuarter < 1 || data.startQuarter > 4) {
      alert("Введите квартал начала периода (целое положительное число от 1 до 4 включительно)!");
      return;
    }

    if (!isNumeric(data.endYear) || data.endYear < 0) {
      alert("Введите год конца периода (целое положительное число)!");
      return;
    }

    if (!isNumeric(data.endQuarter) || data.endQuarter < 1 || data.endQuarter > 4) {
      alert("Введите квартал конца периода (целое положительное число от 1 до 4 включительно)!");
      return;
    }

    if (!isNumeric(data.revenue) || data.revenue < 0) {
      alert("Введите сумму выручки (положительное число)!");
      return;
    }

    if (!isNumeric(data.costs) || data.costs < 0) {
      alert("Введите сумму расходов (положительное число)!");
      return
    }

    console.log(data); 
    alert(JSON.stringify(data));
  };

  return (
    <Form title="Добавить финансовый отчет" buttonText='' fields={financialReportFields} onSubmit={handleSubmit} />
  );
};

export default CreateFinancialReportComponent;
