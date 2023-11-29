import OneCreditImage from '../images/purchase_one_credit.svg'

const PurchaseOneCredit = () => {

  return (
    <div className="flex flex-col vertical-outlined-rectangle bg-FFF9F4">
      1 Request Credit
      <div className="flex items-center justify-center flex-1">
      <img src={OneCreditImage} alt="caits-logo.svg" className="mx-auto" />
      </div>
      <div className='bg-F4E6DC w-full h-59 text-center p-2'>
        $30
      </div>
    </div>
  );
};

export default PurchaseOneCredit;



