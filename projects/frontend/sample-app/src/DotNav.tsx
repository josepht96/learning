import React from 'react';

interface DotNavPanelProps {
    children: React.ReactNode;
    activeIndex: number;
    itemCount: number;
    onClick: (index: number) => void;
}

const DotNavPanel: React.FC<DotNavPanelProps> = ({ activeIndex, itemCount, onClick }) => {
    return (
        <div>
            {Array.from({ length: itemCount }).map((_, index) => (
                <button key={index} onClick={() => onClick(index)} className={activeIndex === index ? 'active' : ''} />
            ))}
        </div>
    );
};

export default DotNavPanel;