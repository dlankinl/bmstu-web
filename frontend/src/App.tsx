import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import CompanyList from './components/Companies/CompanyList';
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  const companies = [
    { name: "Test1", city: "Moscow", description: "Small Test1 Company" },
    { name: "Test2", city: "Moscow", description: "Small Test2 Company" },
    { name: "Test3", city: "Moscow", description: "Small Test3 Company" },
    { name: "Test3", city: "Moscow", description: "Small Test3 Company" },
    { name: "Test3", city: "Moscow", description: "Small Test3 Company" },
    { name: "Test3", city: "Moscow", description: "Small Test3 Company" },
  ]

  return (
    <div className="App">
      <CompanyList companies={companies} />
    </div>
  );
}

export default App
