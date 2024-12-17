import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import CompaniesList from './components/Companies/CompaniesList'
import ActivityFieldsList from './components/ActivityFields/ActivityFieldsList';
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="App">
      {/* <CompaniesList /> */}
      <ActivityFieldsList />
    </div>
  );
}

export default App
