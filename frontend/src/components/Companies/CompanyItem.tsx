import React from 'react';
import { Company } from './types';
import ListItem from '../List/Item';

const CompanyItem: React.FC<{ company: Company }> = ({ company }) => {
  return (
    <ListItem
      item={company}
      renderContent={(item) => (
        <>
          <div className="company-details">
            <div className="company-info">
              <h3>{item.Name}</h3>
              <p>{item.City}</p>
            </div>
            <p className="company-description">{item.Description}</p>
          </div>
        </>
      )}
    />
  );
};

export default CompanyItem;