import React from 'react'
import GiftItem from '../components/GiftItem'
import NavBar from '../components/NavBar'

const gifts = [
  {
    name: 'Gift 1',
    description: 'Gift 1 description',
    price: 1,
    link: 'Gift 1 link',
    demographic: 'Gift 1 demographic',
    giftCollections: 'Gift 1 giftCollections'
  },
  {
    name: 'Gift 2',
    description: 'Gift 2 description',
    price: 2,
    link: 'Gift 2 link',
    demographic: 'Gift 2 demographic',
    giftCollections: 'Gift 2 giftCollections'
  },
  {
    name: 'Gift 3',
    description: 'Gift 3 description',
    price: 3,
    link: 'Gift 3 link',
    demographic: 'Gift 3 demographic',
    giftCollections: 'Gift 3 giftCollections'
  }
]

const GiftManagementPage = () => {
  return (
    <div className='flex flex-col'>
      <NavBar />
      <h1 className='text-2xl text-center'>Gift Management</h1>

      <section className='flex flex-col items-center w-full'>
        <div className='flex flex-row space-x-1 border-b-2'>
          <button className='p-2 text-lg border-r-2 border-r-black'>Gifts</button>
          <button className='p-2 text-lg'>Gift Collections</button>
        </div>
        
        <section className='flex flex-col space-x-1 w-full'>
          
          <button className='border-2 rounded flex flex-col text-center justify-between p-4 max-w-full' style={{margin: '20px 20px'}}>
            <span>+</span>
          </button>

          {gifts.map((gift) => (
            <GiftItem
              key={gift.name}
              name={gift.name}
              description={gift.description}
              price={gift.price}
              link={gift.link}
              demographic={gift.demographic}
              giftCollections={gift.giftCollections}
            />
          ))}
        </section>
      </section>
    </div>
  )
}

export default GiftManagementPage