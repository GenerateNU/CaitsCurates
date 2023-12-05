import React, { useState } from 'react';

interface TextBoxProps {
  placeholder: string;
}

const TextBox: React.FC<TextBoxProps> = ({ placeholder }) => {
  const [searchText, setSearchText] = useState("");

  const handleInputChange = (e: any) => {
    setSearchText(e.target.value);
  };

  return (
    <div className="p-4 w-full">
      <div className="flex items-center relative">
        <input
          type="text"
          style={{ width: '455px', color: '#B79D94', height: '42px', border: '1px solid #B79D94',
          borderWidth: '1.5px'}}
          className="w-full px-4 py-2 rounded-m pl-2"  
          value={searchText}
          onChange={handleInputChange}
          placeholder={placeholder} 
        
        />
      </div>
    </div>
  );
};

export default TextBox;

