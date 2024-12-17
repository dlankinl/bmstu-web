import { useState } from 'react'
import ListContainer from './components/List/ListContainer';
import ActivityFieldsList from './components/ActivityFields/ActivityFieldsList';
import { companyFields, activityFieldFields } from './components/Form/formConfigs';
import Form from './components/Form/Form';
import './App.css'

function App() {
  const [formType, setFormType] = useState<'company' | 'activity field'>('company');

  return (
    <div className="App">
      <ActivityFieldsList isAdmin={true} />
    </div>
  );

  // const handleSubmit = (data: any) => {
  //   console.log(data); // Handle submission logic here
  //   alert(JSON.stringify(data)); // Example of displaying submitted data
  // };

  // return (
  //   <div className="App">
  //     <Form title="Create Company" fields={companyFields} onSubmit={handleSubmit} />
  //   </div>
  // );
}

export default App
