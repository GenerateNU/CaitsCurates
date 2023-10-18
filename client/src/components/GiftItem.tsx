import React from 'react'

interface GiftProps {
  name: string
  description: string
  price: number
  link: string
  demographic: string
  giftCollections: string
}

const GiftItem = (props: GiftProps) => {
  return (
    <div className='border-2 rounded flex flex-col justify-between p-4 max-w-full' style={{margin: '20px 20px'}}>
      <h2>{props.name}</h2>

      <div className='flex flex-col mb-3'>

        {/** some container that holds the prics and link */}
        <div className='flex flex-row space-x-3'>
          <p>Price: ${props.price}</p>
          <a href="#">Buy Now</a>
        </div>
        <p>Demographic: {props.demographic}</p>
        <p>Description: {props.description}</p>
        <p>Collections: {props.giftCollections}</p>
      </div>

      <div className='w-1/12 flex flex-row space-x-2'>
        <button className='px-2 rounded bg-rose-500'>Edit</button>
        <button className='px-1 rounded bg-gray-300'>Delete</button>
      </div>
    </div>
  )
}

export default GiftItem