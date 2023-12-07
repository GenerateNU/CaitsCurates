import React, { useState } from 'react';

const dropdownArrowSVG = `
  <svg width="53" height="46" viewBox="0 0 59 43" xmlns="http://www.w3.org/2000/svg">
    <path d="M0 0.878906H53.9727C56.7341 0.878906 58.9727 3.11748 58.9727 5.87891V37.8184C58.9727 40.5798 56.7341 42.8184 53.9727 42.8184H0V0.878906Z" fill="#F4E6DC"/>
    <path d="M20.4697 19.4742C19.9797 19.9642 19.9797 20.7542 20.4697 21.2442L28.7797 29.5542C29.1697 29.9442 29.7997 29.9442 30.1897 29.5542L38.4997 21.2442C38.9897 20.7542 38.9897 19.9642 38.4997 19.4742C38.0097 18.9842 37.2197 18.9842 36.7297 19.4742L29.4797 26.7142L22.2297 19.4642C21.7497 18.9842 20.9497 18.9842 20.4697 19.4742Z" fill="#BC2C36"/>
  </svg>
`;

interface AgeDropdownProps {
  onChange: (selectedAge: number) => void;
}

const AgeDropdown: React.FC<AgeDropdownProps> = ({ onChange }) => {
  const [selectedAge, setSelectedAge] = useState<number | null>(null);

  const handleAgeChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const age = parseInt(e.target.value, 10);
    setSelectedAge(age);
    onChange(age);
  };

  const dropdownStyle: React.CSSProperties = {
    width: '455px',
    color: '#B79D94',
    height: '40px',
    padding: '4px',
    appearance: 'none',
    margin: '10px',
    marginLeft: '16px',
    marginTop: '16px',
    backgroundImage: `url("data:image/svg+xml,${encodeURIComponent(dropdownArrowSVG)}")`,
    backgroundRepeat: 'no-repeat',
    backgroundPosition: 'right center',
    border: '2px solid #B79D94',
    borderRadius: '5px',
    borderWidth: '1.5px'
  };

  return (
    <select
      value={selectedAge || ''}
      onChange={handleAgeChange}
      style={dropdownStyle}
    >
      <option value="" disabled>
        Age
      </option>
      {Array.from({ length: 100 }, (_, index) => index + 1).map((age) => (
        <option key={age} value={age}>
          {age}
        </option>
      ))}
    </select>
  );
};

export default AgeDropdown;

