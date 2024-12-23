import React, { useEffect, useState } from 'react';
import ListComponent from '../List/List';
import EntrepreneurItem from './EntrepreneurItem';
import { Entrepreneur } from './types';
import ListContainer from '../List/ListContainer';

const BASE_URL = 'http://localhost:8081/api/v2/entrepreneurs';

const EntrepreneursList: React.FC = () => {
  const [entrepreneurs, setEntrepreneurs] = useState<Entrepreneur[]>([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(0); // State for total pages
  const itemsPerPage = 3; // Define how many items per page

  useEffect(() => {
    const fetchEntrepreneurs = async () => {
      try {
        const response = await fetch(`${BASE_URL}?page=${currentPage}&limit=${itemsPerPage}`);
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        setEntrepreneurs(data.entrepreneurs); // Assuming data.entrepreneurs contains an array of entrepreneurs
        setTotalPages(data.numPages-1); // Set total pages from API response
      } catch (error) {
        console.error("Error fetching entrepreneurs:", error);
      }
    };

    fetchEntrepreneurs();
  }, [currentPage]); // Fetch data whenever currentPage changes

  return (
    <ListContainer>
      <div style={{ position: 'relative' }}>
        <ListComponent
          data={entrepreneurs}
          renderItem={(entrepreneur) => (
            <EntrepreneurItem key={entrepreneur.id} entrepreneur={entrepreneur} />
          )}
          itemsPerPage={itemsPerPage}
          currentPage={currentPage} // Pass current page to ListComponent
          totalPages={totalPages} // Pass total pages to ListComponent
          onPageChange={setCurrentPage} // Pass down the page change handler
        />
      </div>
    </ListContainer>
  );
};

export default EntrepreneursList;
