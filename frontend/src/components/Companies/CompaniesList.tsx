// import React, { useEffect, useState } from 'react';
// import ListComponent from '../List/List';
// import CompanyItem from './CompanyItem';
// import { Company } from './types';
// import ListContainer from '../List/ListContainer';
// import GreenButton from '../Buttons/GreenButton';
// import "./CompaniesList.css";
// import CreateCompanyComponent from '../CreateCompany/CreateCompany';
// import Modal from '../Modal/Modal';

// const CompaniesList: React.FC = () => {
//   const [isModalActive, setModalActive] = useState(false);

//   const handleModalOpen = () => {
//     setModalActive(true);
//   };
//   const handleModalClose = () => {
//     setModalActive(false);
//   };

//   const [companies, setCompanies] = useState<Company[]>([]);

//   const isOwner = true;

//   useEffect(() => {
//     // Simulate fetching data
//     const fetchCompanies = () => {
//       const staticCompanies: Company[] = [
//         { ID: "1", Name: "Test1", City: "Moscow", Description: "Small Test1 Company", OwnerID: "owner1", ActivityFieldID: "AF1" },
//         { ID: "2", Name: "Test2", City: "Moscow", Description: "Small Test2 Company", OwnerID: "owner2", ActivityFieldID: "AF2" },
//         { ID: "3", Name: "Test3", City: "Moscow", Description: "Small Test3 Company", OwnerID: "owner3", ActivityFieldID: "AF3" },
//         { ID: "4", Name: "Test4", City: "Moscow", Description: "Small Test4 Company", OwnerID: "owner4", ActivityFieldID: "AF4" },
//         { ID: "5", Name: "Test5", City: "Moscow", Description: "Small Test5 Company", OwnerID: "owner5", ActivityFieldID: "AF5" },
//         { ID: "6", Name: "Test6", City: "Moscow", Description: "Small Test6 Company", OwnerID: "owner6", ActivityFieldID: "AF6" },
//       ];
//       setCompanies(staticCompanies);
//     };

//     fetchCompanies();
//   }, []);

//   return (
//     <ListContainer>
//       <div style={{ position: 'relative' }}>
//         {isOwner && (
//           <div className="button-container">
//             <GreenButton text={"Добавить"} icon={""} onClick={handleModalOpen}/>
//           </div>
//         )}
//         <ListComponent
//           data={companies}
//           renderItem={(company) => (
//             <CompanyItem key={company.ID} company={company} />
//           )}
//           itemsPerPage={3}
//         />

//         <div>
//           {isModalActive && (
//             <Modal onClose={handleModalClose}>
//               <CreateCompanyComponent isEditing={false}/>
//             </Modal>
//           )}
//         </div>
//       </div>
//     </ListContainer>
//   );
// };

// export default CompaniesList;


import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import ListComponent from '../List/List';
import CompanyItem from './CompanyItem';
import { Company } from './types';
import ListContainer from '../List/ListContainer';
import GreenButton from '../Buttons/GreenButton';
import "./CompaniesList.css";
import CreateCompanyComponent from '../CreateCompany/CreateCompany';
import Modal from '../Modal/Modal';

const BASE_URL = 'http://localhost:8081/api/v2/entrepreneurs'; // Adjust this URL to your API endpoint

const CompaniesList: React.FC = () => {
  const [isModalActive, setModalActive] = useState(false);
  const [companies, setCompanies] = useState<Company[]>([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(0); // State for total pages

  const handleModalOpen = () => {
    setModalActive(true);
  };

  const handleModalClose = () => {
    setModalActive(false);
  };

  const { id } = useParams<{ id: string }>();

  var isOwner = false;
  const parts = localStorage.getItem("authToken")?.split('.');
  if (parts !== undefined) {
    const decodedPayload = atob(parts[1]);

    const parsedPayload = JSON.parse(decodedPayload);
    isOwner = parsedPayload.sub === id;
  }

  useEffect(() => {
    const fetchCompanies = async () => {
      try {
        const response = await fetch(`${BASE_URL}/${id}/companies?page=${currentPage}`); // Adjust API call as needed
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        setCompanies(data.companies);
        setTotalPages(data.numPages);
      } catch (error) {
        console.error("Error fetching companies:", error);
      }
    };

    fetchCompanies();
  }, [currentPage]); // Fetch data whenever currentPage changes

  return (
    <ListContainer>
      <div style={{ position: 'relative' }}>
        {isOwner && (
          <div className="button-container">
            <GreenButton text={"Добавить"} icon={""} onClick={handleModalOpen}/>
          </div>
        )}
        
        <ListComponent
          data={companies}
          renderItem={(company) => (
            <CompanyItem key={company.id} company={company} />
          )}
          itemsPerPage={3}
          currentPage={currentPage} // Pass current page to ListComponent
          totalPages={totalPages} // Pass total pages to ListComponent
          onPageChange={setCurrentPage} // Pass down the page change handler
        />

        <div>
          {isModalActive && (
            <Modal onClose={handleModalClose}>
              <CreateCompanyComponent isEditing={false}/>
            </Modal>
          )}
        </div>
      </div>
    </ListContainer>
  );
};

export default CompaniesList;
