import React, { useState } from 'react';
import ErrorPopup from '../ErrorPopup/ErrorPopUp';
import { financialReportFormFields } from '../Form/formConfigs';
import Form from '../Form/Form';

interface FinancialReportFormProps {
  onSubmit: (data: { startYear: number; endYear: number; startQuarter: number; endQuarter: number }) => void;
}

// const FinancialReportForm: React.FC<FinancialReportFormProps> = ({ onSubmit }) => {
//   const [startYear, setStartYear] = useState<number>(0);
//   const [endYear, setEndYear] = useState<number>(0);
//   const [startQuarter, setStartQuarter] = useState<number>(1);
//   const [endQuarter, setEndQuarter] = useState<number>(1);

//   const handleSubmit = (e: React.FormEvent) => {
//     e.preventDefault();
//     onSubmit({ startYear, endYear, startQuarter, endQuarter });
//   };

//   return (
//     <FormContainer>
//       <form onSubmit={handleSubmit}>
//         <div>
//           <label>Год начала:</label>
//           <input type="number" value={startYear} onChange={(e) => setStartYear(Number(e.target.value))} required />
//         </div>
//         <div>
//           <label>Год конца:</label>
//           <input type="number" value={endYear} onChange={(e) => setEndYear(Number(e.target.value))} required />
//         </div>
//         <div>
//           <label>Квартал начала:</label>
//           <input type="number" value={startQuarter} onChange={(e) => setStartQuarter(Number(e.target.value))} min={1} max={4} required />
//         </div>
//         <div>
//           <label>Квартал конца:</label>
//           <input type="number" value={endQuarter} onChange={(e) => setEndQuarter(Number(e.target.value))} min={1} max={4} required />
//         </div>
//         <button type="submit">Получить отчет</button>
//       </form>
//     </FormContainer>
//   );
// };

// export default FinancialReportForm;

function isNumeric(value: any) {
  return /^\d+(\.\d+)?$/.test(value);
}

const FinancialReportForm: React.FC<FinancialReportFormProps> = ({ onSubmit }) => {
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [showErrorPopup, setShowErrorPopup] = useState(false);

  const handleCloseError = () => {
    setShowErrorPopup(false);
    setErrorMessage(null);
  };

  const handleSubmit = (data: any) => {
    if (!isNumeric(data.startYear) || data.startYear < 0) {
      setErrorMessage("Введите год начала периода (целое положительное число)!");
      setShowErrorPopup(true);
      return;
    }

    if (!isNumeric(data.startQuarter) || data.startQuarter < 1 || data.startQuarter > 4) {
      setErrorMessage("Введите квартал начала периода (целое положительное число от 1 до 4 включительно)!");
      setShowErrorPopup(true);
      return;
    }

    if (!isNumeric(data.endYear) || data.endYear < 0) {
      setErrorMessage("Введите год конца периода (целое положительное число)!");
      setShowErrorPopup(true);
      return;
    }

    if (data.endYear < data.startYear) {
      setErrorMessage("Год конца периода должен быть больше либо равен году начала!");
      setShowErrorPopup(true);
      return;
    }

    if (!isNumeric(data.endQuarter) || data.endQuarter < 1 || data.endQuarter > 4) {
      setErrorMessage("Введите квартал конца периода (целое положительное число от 1 до 4 включительно)!");
      setShowErrorPopup(true);
      return;
    }

    console.log(data); 
    alert(JSON.stringify(data));
    onSubmit(data);
  };

  return (
    <>
      <Form title="Укажите период" buttonText='' fields={financialReportFormFields} onSubmit={handleSubmit} />
      {showErrorPopup && errorMessage && (
        <ErrorPopup message={errorMessage} onClose={handleCloseError} />
      )}
    </>
  );
};

export default FinancialReportForm;
