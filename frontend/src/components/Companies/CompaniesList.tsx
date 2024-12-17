import React, { useEffect, useState } from 'react';
// import ListComponent from '../ListComponent';
import ListComponent from '../List/List';
import CompanyItem from './CompanyItem';
import { Company } from './types';

const CompaniesList: React.FC = () => {
  const [companies, setCompanies] = useState<Company[]>([]);

  useEffect(() => {
    // Simulate fetching data
    const fetchCompanies = () => {
      const staticCompanies: Company[] = [
        { ID: "1", Name: "Test1", City: "Moscow", Description: "Small Test1 Company", OwnerID: "owner1", ActivityFieldID: "AF1" },
        { ID: "2", Name: "Test2", City: "Moscow", Description: "Small Test2 Company", OwnerID: "owner2", ActivityFieldID: "AF2" },
        { ID: "3", Name: "Test3", City: "Moscow", Description: "Small Test3 Company", OwnerID: "owner3", ActivityFieldID: "AF3" },
        { ID: "4", Name: "Test4", City: "Moscow", Description: "Small Test4 Company", OwnerID: "owner4", ActivityFieldID: "AF4" },
        { ID: "5", Name: "Test5", City: "Moscow", Description: "Small Test5 Company", OwnerID: "owner5", ActivityFieldID: "AF5" },
        { ID: "6", Name: "Test6", City: "Moscow", Description: "Small Test6 Company", OwnerID: "owner6", ActivityFieldID: "AF6" },
      ];
      setCompanies(staticCompanies);
    };

    fetchCompanies();
  }, []);

  return (
    <ListComponent
      data={companies}
      renderItem={(company) => (
        <CompanyItem key={company.ID} company={company} />
      )}
      itemsPerPage={3}
    />
  );
};

export default CompaniesList;