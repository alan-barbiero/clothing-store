import React from 'react'
import AddClothing from './components/AddClothing'

function App() {
  return (
    <div className='container' style={{display:'flex', justifyContent: 'center', flexDirection: 'column', alignItems: 'center',  gap: '2rem', height: '60vh'}}>
      <h1>Register Clothes</h1>
      <AddClothing />
    </div>
  )
}

export default App