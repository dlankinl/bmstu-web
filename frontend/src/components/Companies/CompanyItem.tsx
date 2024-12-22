import React from 'react';
import { Company } from './types';
import ListItem from '../List/Item';
import { Link } from 'react-router-dom';

const CompanyItem: React.FC<{ company: Company }> = ({ company }) => {
  return (
    <ListItem
      item={company}
      renderContent={(item) => (
        <>
          <div className="item-details">
            <div className="item-info">
              <Link to={`/companies/${item.ID}`}>
                <h3>{item.Name}</h3>
              </Link>
              <p>{item.City}</p>
            </div>
            <p className="item-description">{item.Description}</p>
          </div>
        </>
      )}
    />
  );
};

export default CompanyItem;