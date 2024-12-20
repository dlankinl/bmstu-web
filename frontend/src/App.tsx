import { useState } from 'react'
import ListContainer from './components/List/ListContainer';
import ActivityFieldsList from './components/ActivityFields/ActivityFieldsList';
import { companyFields, activityFieldFields } from './components/Form/formConfigs';
import Form from './components/Form/Form';
import Navbar from './components/Navbar/NavBar';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import CompanyInfoComponent from './components/CompanyInfo/CompanyInfo';
import './App.css'

function App() {
  const [formType, setFormType] = useState<'company' | 'activity field'>('company');

  const handleSubmit = (data: any) => {
    console.log(data); // Handle submission logic here
    alert(JSON.stringify(data)); // Example of displaying submitted data
  };

  return (
    <Router>
      <div className="App">
        <Navbar /> {/* Insert Navbar here */}
        <Routes>
          {/* <Route path="/" element={<Form title="Создать компанию" fields={companyFields} onSubmit={handleSubmit} />} /> */}
          {/* <Route path="/" element={<CompanyInfo values={["val1", "val2", "val3", "val4"]}/>} /> */}
          <Route path="/companies/:id" element={<CompanyInfoComponent />} />
          <Route path="/activity-fields" element={<ActivityFieldsList isAdmin={true} />} />
          <Route path="/companies/create" element={<Form title="Создать компанию" fields={companyFields} onSubmit={handleSubmit} />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App
