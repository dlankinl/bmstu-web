import { useState } from 'react'
import ListContainer from './components/List/ListContainer';
import ActivityFieldsList from './components/ActivityFields/ActivityFieldsList';
// import { companyFields, activityFieldFields } from './components/Form/formConfigs';
import Form from './components/Form/Form';
import Navbar from './components/Navbar/Navbar';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import CompanyInfoComponent from './components/CompanyInfo/CompanyInfo';
import './App.css'
import CompaniesList from './components/Companies/CompaniesList';
import FinancialReportComponent from './components/FinancialReport/FinancialReport';
import SignupComponent from './components/Signup/Signup';
import EntrepreneursList from './components/Entrepreneurs/EntrepreneursList';
import EntrepreneurInfoComponent from './components/EntrepreneurInfo/EntrepreneurInfo';

function App() {
  // const [formType, setFormType] = useState<'company' | 'activity field'>('company');

  // const handleSubmit = (data: any) => {
  //   console.log(data); 
  //   alert(JSON.stringify(data));
  // };

  return (
    <Router>
      <div className="App">
        <Navbar />
        <Routes>
          <Route path="/companies/:id" element={<CompanyInfoComponent />} />
          <Route path="/entrepreneurs/:id" element={<EntrepreneurInfoComponent />} />
          <Route path="/activity-fields" element={<ActivityFieldsList isAdmin={true} />} />
          <Route path="/companies" element={<CompaniesList />} />
          {/* <Route path="/companies/create" element={<Form title="Создать компанию" fields={companyFields} onSubmit={handleSubmit} />} /> */}
          <Route path="/financials/:id" element={<FinancialReportComponent />} />
          {/* <Route path="/signup" element={<SignupComponent />} /> */}
          <Route path="/entrepreneurs" element={<EntrepreneursList />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App
