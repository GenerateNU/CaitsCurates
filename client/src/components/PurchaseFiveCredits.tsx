import FiveCreditImage from '../images/purchase_five_credits.svg'

const PurchaseFiveCredits = () => {

    return (
      <div className="flex flex-col vertical-outlined-rectangle bg-FFF9F4">
        5 Request Credits
        <div className="flex items-center justify-center flex-1">
        <img src={FiveCreditImage} alt="caits-logo.svg" className="mx-auto" />
        </div>
        <div className='bg-F4E6DC w-full h-59 text-center p-2'>
          $55
        </div>
      </div>
    );
  };
  
  export default PurchaseFiveCredits;
  
  
  
  