import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import axios from 'axios'; // Import axios

function App() {
  const [count, setCount] = useState(0);
  const [backendData, setBackendData] = useState(null); // State to store data from backend
  const [error, setError] = useState(null); // State to store any errors

  // Function to fetch data from the backend
  const fetchBackendData = async () => {
    try {
      // *** Replace with your actual backend URL and endpoint ***
      const response = await axios.get('http://localhost:3000/api/data');
      console.log('Data from backend:', response.data);
      setBackendData(response.data); // Update state with fetched data
      setError(null); // Clear any previous errors
    } catch (err) {
      console.error('Error fetching data:', err);
      setError(err); // Store the error
      setBackendData(null); // Clear any previous data
    }
  };

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.jsx</code> and save to test HMR
        </p>
      </div>
      {/* New button to fetch data from backend */}
      <div className="card">
        <button onClick={fetchBackendData}>
          Fetch Data from Backend
        </button>
        {/* Display backend data or error */}
        {backendData && (
          <p>Backend Data: {JSON.stringify(backendData)}</p> // Display data (adjust as needed)
        )}
        {error && (
          <p style={{ color: 'red' }}>Error: {error.message}</p>
        )}
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
