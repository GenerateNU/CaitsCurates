import React, { useState } from 'react';

interface DateTextBoxProps {
    onChange: (value: string) => void;
    width: string;
  }

const DateTextBox: React.FC<DateTextBoxProps> = ({onChange, width}) => {

    const [searchDate, setSearchDate] = useState('');

    const handleInputChange = (e: any) => {
        setSearchDate(e.target.value);
        onChange(e.target.value);
    };
  
    return (
        <div className="flex items-center relative">
            <input
            type="text"
            style={{
                color: '#B79D94',
                width: width,
                height: '42px',
                border: '1.5px solid #B79D94',
                borderRadius: '5px',
                borderWidth: '1.5px'
            }}
            className="w-full px-4 py-2 rounded-m pl-2"
            value={searchDate}
            placeholder={"Date"}
            onChange={handleInputChange}
            />
        </div>
    );
  };
  
  export default DateTextBox;
  