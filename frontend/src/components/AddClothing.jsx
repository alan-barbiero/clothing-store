import { useState } from 'react'

export default function AddClothing({ onAdd }) {
  const [name, setName] = useState('')
  const [size, setSize] = useState('')
  const [category, setCategory] = useState('')
  const [quantity, setQuantity] = useState('')
  const [price, setPrice] = useState('')
  const [error, setError] = useState('')

  const handleSubmit = async (e) => {
    e.preventDefault()
    setError('')

    const requestData = {
      name: name.trim(),
      size: size.trim(),
      category: category.trim(),
      quantity: parseInt(quantity),
      price: parseFloat(price)
    }

    try {
      const response = await fetch('http://localhost:6688/api/clothing', {
        method: 'POST',
        mode: 'cors',
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json'
        },
        body: JSON.stringify(requestData)
      })

      if (!response.ok) {
        const errorText = await response.text()
        throw new Error(`HTTP error! status: ${response.status}, message: ${errorText}`)
      }

      const data = await response.json()
      alert('Successfully registered!')
      console.log('Response:', data)

      // Clear form
      setName('')
      setSize('')
      setPrice('')
      setCategory('')
      setQuantity('')

      // Call onAdd if it exists
      if (typeof onAdd === 'function') {
        onAdd()
      }
    } catch (err) {
      console.error('Registration error:', err)
      setError(err.message)
      alert(`Error: ${err.message}`)
    }
  }

  return (
    <div className='container' style={{display:'flex', justifyContent:'center',}} >
    <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', gap: '1rem', maxWidth: '300px' }}>
      {error && <div style={{ color: 'red' }}>{error}</div>}
      <input type="text" placeholder="Name" value={name} onChange={(e) => setName(e.target.value)} required />
      <input type="text" placeholder="Size" value={size} onChange={(e) => setSize(e.target.value)} required />
      <input type="text" placeholder="Category" value={category} onChange={(e) => setCategory(e.target.value)} required />
      <input type="number" step="1" placeholder="Quantity" value={quantity} onChange={(e) => setQuantity(e.target.value)} required />
      <input type="number" step="0.01" placeholder="Price" value={price} onChange={(e) => setPrice(e.target.value)} required />
      <button type="submit">Register Clothing</button>
    </form>
    </div>
  )
}
