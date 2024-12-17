import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import ListContainer from './components/List/ListContainer';
import ActivityFieldsList from './components/ActivityFields/ActivityFieldsList';
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="App">
      <ListContainer>
        <ActivityFieldsList />
      </ListContainer>
    </div>
  );
}

export default App
