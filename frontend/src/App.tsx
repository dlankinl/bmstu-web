import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
// import CompanyList from './components/Companies/CompanyList';
import CompaniesList from './components/Companies/CompaniesList'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="App">
      {/* <CompanyList /> */}
      <CompaniesList />
    </div>
  );
}

export default App
