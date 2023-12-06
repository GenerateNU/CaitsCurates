import React, { useState } from 'react';

interface BudgetSliderProps {
    maxBudget: number;
    minBudget: number;
  }

const BudgetSlider: React.FC<BudgetSliderProps> = ({maxBudget, minBudget}) => {

    const [sliderValue, setSliderValue] = useState(250);

    const handleSliderChange = (e: any) => {
        setSliderValue(e.target.value);
    };
  
    return (
        <div className="flex items-center relative">
            <input
            type="range"
            min={minBudget}
            max={maxBudget}
            step="1"
            style={{
                background: '#B79D94',
                width: '455px',
                marginLeft: '18px',
                marginTop: '12px',
                height: '42px',
                border: '1.5px solid #B79D94',
                borderRadius: '5px',
                borderWidth: '1.5px'
            }}
            className="w-full rounded-m"
            value={sliderValue}
            onChange={handleSliderChange}
            />
        </div>
    );
  };
  
  export default BudgetSlider;
  