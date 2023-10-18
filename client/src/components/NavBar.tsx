import React from 'react'

const NavBar = () => {
  return (
    <div className='bg-rose-500 flex flex-row justify-between items-center p-2 w-full'>
        <h1>Caits Curates</h1>

        <div className='flex flex-row space-x-1'>
            <button className='border rounded flex flex-col items-center p-2'>
            +
            </button>
            <button className='border rounded flex flex-col items-center p-2'>
            -
            </button>
        </div>
    </div>
  )
}

export default NavBar