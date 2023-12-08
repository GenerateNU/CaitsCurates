import React, {useState} from 'react';
import ModalScreenFour from './ModalScreenFour.tsx';
import ModalScreenFinal from './ModalFinal.tsx';
import ModalX from '../../images/modal_x_out.svg';
import { GiftRequestProps} from '../../types.tsx'
import axios from 'axios';

interface ModalScreenFiveProps {
    isOpen: boolean;
    onClose: () => void;
    children: React.ReactNode;
    gifteeID: number;
    giftRequest: GiftRequestProps;
}

const ModalScreenFive: React.FC<ModalScreenFiveProps> = ({ isOpen, onClose, children, giftRequest, gifteeID }) => {
    const [showModalFour, setShowModalFour] = useState(false);
    const [showModalFinal, setShowModalFinal] = useState(false);
    const [comment, setComment] = useState("");
    const handleNextClick = () => {
        makeRequest()
        updateAvailableRequests()
        setShowModalFinal(true);
    };

    const handleBackClick = () => {
        setShowModalFour(true);
    };

    // Endpoint Calls
    const makeRequest = async () => {
        try {
            const response = await axios.post("/api/addGiftRequest", {
                BudgetMax: giftRequest.BudgetMax,
                BudgetMin: giftRequest.BudgetMin,
                Comment: comment,
                CustomerID: giftRequest.CustomerID,
                DateNeeded: giftRequest.DateNeeded,
                GifteeID: gifteeID,
                Occasion: []
            });
            return response.data;
        } catch (error) {
            console.error("An error occured while making the request");
        }
    };
    const updateAvailableRequests = async () => {
        try {
            const response = await axios.put(`/api/customer/1/-1`);
            return response.data;
        } catch (error) {
            console.log("An error when occured while updating Available Requests");
        }
    };

    const handleCommentChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setComment(event.target.value);
    };


    const overlayStyles = isOpen
        ? 'fixed top-0 left-0 w-full h-full flex justify-center items-center z-50'
        : 'hidden';

    const contentStyles = isOpen
        ? 'bg-FFF9F4 rounded-md shadow-md flex flex-col px-50 items-center'
        : 'hidden';

    const buttonStyle = "px-4 py-2 ml-40 text-FFF9F4 text-s bg-273F2A rounded-md";
    const activeButtonStyle = "bg-273F2A";

    return (
        <div className={overlayStyles} onClick={onClose}>
            <div className={contentStyles} onClick={(e) => e.stopPropagation()}>
                {showModalFour ? (
                    <ModalScreenFour isOpen={isOpen} onClose={onClose} children={children} gifteeID={gifteeID}/>
                ) : (
                    <>
                        {showModalFinal ? (
                            <ModalScreenFinal isOpen={isOpen} onClose={onClose} children={children} date = {giftRequest.DateNeeded} />
                        ) : (
                            <>
                                <div className="flex ml-10 items-center mt-10">
                                    <button className=''
                                            onClick={() => handleBackClick()}
                                    >
                                        Back
                                    </button>
                                    <div className="mr-auto ml-56">
                                        <img
                                            src={ModalX}
                                            alt="caits-logo.svg"
                                            onClick={onClose}
                                            style={{ height: "20px", width: "20px", marginLeft: "880px", marginRight: "20px"}}
                                        />
                                    </div>
                                </div>
                                <div>
                                    <h1 className="font-bold mb-10 text-3xl text-center items-center w-full font-seasons">Giftee Preferences</h1>

                                    <div className= "flex flex-col mb-14">
                                        <h2 className='text-2xl text-A65A5A font-seasons font-bold leading-7'>Any questions or comments?</h2>
                                        <textarea
                                            style={{
                                                width: '740px',
                                                height: '204.52px',
                                                border: '1.5px solid #B79D94',
                                                borderRadius: '5px',
                                                boxSizing: 'border-box',
                                                padding: '12px',
                                                fontSize: '16px',
                                                resize: 'none',
                                            }}
                                            value ={comment}
                                            onChange={handleCommentChange}
                                            className='mt-6 text-B79D94 font-proxima mr-24'
                                            placeholder='Questions, comments, concerns, or anything else I should know?'
                                        />
                                    </div>
                                    <div className='mb-10 flex justify-end'>
                                        <button
                                            className={`${buttonStyle} ${activeButtonStyle}`}
                                            onClick={() => handleNextClick()}
                                            style={{ width: "170px", height: "50px"}}
                                        >
                                            Submit
                                        </button>
                                    </div>
                                </div>
                            </>
                        )}
                    </>
                )}
            </div>
        </div>
    );
};

export default ModalScreenFive;