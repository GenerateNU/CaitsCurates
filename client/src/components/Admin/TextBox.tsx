import React, { useState } from 'react';

interface TextBoxProps {
    placeholder: string;
    onChange: (value: string) => void;
  }

const TextBox: React.FC<TextBoxProps> = ({ placeholder, onChange }) => {

    const [searchText, setSearchText] = useState('');

    const handleInputChange = (e: any) => {
        setSearchText(e.target.value);
        onChange(e.target.value);
    };
  
    return (
      <div className="p-4 w-full">
        <div className="flex items-center relative">
          <input
            type="text"
            style={{
              width: '455px',
              color: '#B79D94',
              height: '42px',
              border: '1.5px solid #B79D94',
              borderRadius: '5px',
              borderWidth: '1.5px'
            }}
            className="w-full px-4 py-2 rounded-m pl-2"
            value={searchText}
            placeholder={placeholder}
            onChange={handleInputChange}
          />
        </div>
      </div>
    );
  };
  
  export default TextBox;
  