import React, { useState, useEffect } from 'react';
import { Company } from './types';
import Pagination from '../Pagination/Pagination';

interface CompanyItemProps {
  company: Company;
}

const CompanyItem: React.FC<CompanyItemProps> = ({ company }) => {
  return (
    <div className="company-item">
      <div className="company-details">
        <div className="company-info">
          <h3>{company.Name}</h3>
          <p>{company.City}</p>
        </div>
        <p className="company-description">{company.Description}</p>
      </div>
    </div>
  );
};

const CompanyList: React.FC = () => {
  const [companies, setCompanies] = useState<Company[]>([]);
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 3;
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const staticCompanies: Company[] = [
    { ID: 1, OwnerID: 101, Name: "Test1", City: "Moscow", Description: "Small Test1 Company", ActivityFieldID: 1 },
    { ID: 2, OwnerID: 102, Name: "Test2", City: "Moscow", Description: "Small Test2 Company", ActivityFieldID: 2 },
    { ID: 3, OwnerID: 103, Name: "Test3", City: "Moscow", Description: "Small Test3 Company", ActivityFieldID: 3 },
    { ID: 4, OwnerID: 104, Name: "Test4", City: "Moscow", Description: "Small Test4 Company", ActivityFieldID: 4 },
    { ID: 5, OwnerID: 105, Name: "Test5", City: "Moscow", Description: "Small Test5 Company", ActivityFieldID: 5 },
    { ID: 6, OwnerID: 106, Name: "Test6", City: "Moscow", Description: "Small Test6 Company", ActivityFieldID: 6 },
  ];

  useEffect(() => {
    // const fetchCompanies = async () => {
    //   try {
    //     const response = await fetch('https://api.example.com/companies'); // Replace with your API endpoint
    //     if (!response.ok) {
    //       throw new Error('Network response was not ok');
    //     }
    //     const data = await response.json();
    //     setCompanies(data);
    //   } catch (error) {
    //     setError(error.message);
    //   } finally {
    //     setLoading(false);
    //   }
    // };
    const fetchCompanies = () => {
      setCompanies(staticCompanies);
      setLoading(false);
    }

    fetchCompanies();
  }, []);

  const indexOfLastItem = currentPage * itemsPerPage;
  const indexOfFirstItem = indexOfLastItem - itemsPerPage;
  const currentCompanies = companies.slice(indexOfFirstItem, indexOfLastItem);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;

  return (
    <div className="company-list">
      <h2>Companies</h2>
      {currentCompanies.map((company) => (
        <CompanyItem key={company.ID} company={company} /> // Assuming each company has a unique 'id'
      ))}
      <Pagination
        currentPage={currentPage}
        totalPages={Math.ceil(companies.length / itemsPerPage)}
        onPageChange={setCurrentPage}
      />
    </div>
  );
};

export default CompanyList;