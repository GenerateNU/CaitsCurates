import {Gift} from "../../types.tsx";



interface GiftProps {
    gift: Gift;
    onEditClick: (gift: Gift) => void;
    onDeleteClick: (gift: Gift) => void;

}

const GiftItem = (props: GiftProps) => {
    const giftElements = props.gift?.Category?.map((category) => (
        <span key={category} className="px-2 py-1 bg-gray-200 rounded-full">{category}</span>
    ));
    return (
        <div className='border-2 rounded flex flex-col justify-between p-4 max-w-full' style={{ margin: '20px 20px' }}>
            <h2>{props.gift.Name}</h2>

            <div className='flex flex-col mb-3'>
                <div className='flex flex-row space-x-3'>
                    <p>Price: ${props.gift.Price}</p>
                    <a href={props.gift.Link}
                       className="text-blue-500 hover:underline hover:text-blue-700"
                    >Link</a>
                </div>
                <p>Demographic: {props.gift.Demographic}</p>
                <p>Categories:</p>
                <div className='flex flex-wrap gap-1'>
                    {giftElements}
                </div>
                <p>Description: {props.gift.Description}</p>
                <p>Occasion: {props.gift.Occasion}</p>
                <p>Image Link: <a href={props.gift.ImageLink}>link</a></p>
            </div>

                <div className='w-1/12 flex flex-row space-x-2'>
                    <button
                        className='px-2 rounded bg-blue-400'
                        onClick={() => props.onEditClick(props.gift)}
                    >
                        Edit
                    </button>
                    <button className='px-1 rounded bg-gray-300'
                            onClick={() => props.onDeleteClick(props.gift)}>Delete</button>
                </div>
            </div>
    )
}

export default GiftItem;