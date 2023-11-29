import ThreeCreditImage from '../images/purchase_three_credit.svg'

const PurchaseThreeCredits = () => {

    return (
      <div className="flex flex-col vertical-outlined-rectangle bg-FFF9F4">
        3 Request Credits
        <div className="flex items-center justify-center flex-1">
        <img src={ThreeCreditImage} alt="caits-logo.svg" className="mx-auto" />
        </div>
        <div className='bg-F4E6DC w-full h-59 text-center p-2'>
          $45
        </div>
      </div>
      
    );
  };
  
  export default PurchaseThreeCredits;
  
  
  


