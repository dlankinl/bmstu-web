import React from 'react';
import { useParams } from 'react-router-dom';
import { useState, useEffect } from 'react';
import './FinancialReport.css';
import GreenButton from '../Buttons/GreenButton';
import FormContainer from '../Form/FormContainer';
import { FinancialReport } from './types';

interface FinancialReportParamsProps {
  startYear: number;
  startQuarter: number;
  endYear: number;
  endQuarter: number;
  companyID: number | null;
  entrepreneurID: number | null;
}

const FinancialReportComponent: React.FC<FinancialReportParamsProps> = ({ params }) => {
  const { id } = useParams();
  const [value, setValues] = useState<FinancialReport>();

  useEffect(() => {
    const fetchFinancialReport = () => {
      const staticFinancialReport: FinancialReport = { ID: "1", CompanyID: "1", Revenue: 123, Costs: 23, Profit: 100, 
        PeriodStart: `${params.startQuarter}Q ${params.startYear}`, PeriodEnd: `${params.endQuarter}Q ${params.endYear}`, CompanyName: "test1" };
      setValues(staticFinancialReport);
    };

    // const fetchData = async () => {
    //   try {
    //     const response = await fetch(`https://api.example.com/companies/${id}`);
    //     const data = await response.json();
    //     setValues(data);
    //   } catch (error) {
    //     console.error("Error fetching company data:", error);
    //   }
    // };

    fetchFinancialReport();
  }, [id]); 

  const label = !params.companyID && params.entrepreneurID ? "ФИО" : "Объект";

  return (
    <FormContainer>
      <div className="my-component">
        <div className="row">
          <span>{label}</span>
          <span>{value?.CompanyName}</span>
        </div>
        <div className="row">
          <span>Начало периода</span>
          <span>{value?.PeriodStart}</span>
        </div>
        <div className="row">
          <span>Конец периода</span>
          <span>{value?.PeriodEnd}</span>
        </div>
        <div style={{ height: '25px' }}></div>
        <div className="row">
          <span>Выручка</span>
          <span>{value?.Revenue}</span>
        </div>
        <div className="row">
          <span>Расходы</span>
          <span>{value?.Costs}</span>
        </div>
        <div className="row">
          <span>Прибыль</span>
          <span>{value?.Profit}</span>
        </div>
      </div>
    </FormContainer>
  );
};

export default FinancialReportComponent;
