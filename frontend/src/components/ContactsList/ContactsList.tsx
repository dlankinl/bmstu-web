import React from 'react';
import { useParams } from 'react-router-dom';
import { useState, useEffect } from 'react';
import './ContactsList.css';
import FormContainer from '../Form/FormContainer';
import { ContactInfo } from './types';

const ContactsList = ({ }) => {
  const { id } = useParams();
  const [values, setValues] = useState<ContactInfo[]>([]);

  useEffect(() => {
    const fetchContactsInfo = () => {
      const staticContactsInfo: ContactInfo[] = [
        { ID: "1", Name: "telegram", Value: "durov"},
        { ID: "2", Name: "ok", Value: "durov"},
      ];
      setValues(staticContactsInfo);
    };

    // const fetchData = async () => {
    //   try {
    //     const response = await fetch(`https://api.example.com/companies/${id}`);
    //     const data = await response.json();
    //     setValues(data);
    //   } catch (error) {
    //     console.error("Error fetching company data:", error);
    //   }
    // };

    fetchContactsInfo();
  }, [id]); 

  return (
    <FormContainer>
      <div className="my-component">
        {values.map(contact => (
          <div key={contact.ID} className="row">
            <span>{contact.Name}</span>
            <span>{contact.Value}</span>
          </div>
        ))}
      </div>
    </FormContainer>
  );
};

export default ContactsList;
