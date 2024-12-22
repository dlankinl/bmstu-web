import React, { useEffect, useState } from 'react';
import ListComponent from '../List/List';
import EntrepreneurItem from './EntrepreneurItem';
import { Entrepreneur } from './types';
import ListContainer from '../List/ListContainer';

const EntrepreneursList: React.FC = () => {
  const [entrepreneurs, setEntrepreneurs] = useState<Entrepreneur[]>([]);

  useEffect(() => {
    const fetchEntrepreneurs = () => {
      const staticEntrepreneurs: Entrepreneur[] = [
        { ID: "1", Name: "Test1", City: "Moscow", Birthday: "22 декабря 2000 г.", Rating: 5.0, Gender: "мужской" },
        { ID: "2", Name: "Test2", City: "Moscow", Birthday: "23 декабря 2000 г.", Rating: 4.0, Gender: "женский" },
        { ID: "3", Name: "Test3", City: "Moscow", Birthday: "24 декабря 2000 г.", Rating: 3.0, Gender: "мужской" },
        { ID: "4", Name: "Test4", City: "Moscow", Birthday: "25 декабря 2000 г.", Rating: 3.0, Gender: "мужской" },
        { ID: "5", Name: "Test5", City: "Moscow", Birthday: "26 декабря 2000 г.", Rating: 4.0, Gender: "женский" },
        { ID: "6", Name: "Test6", City: "Moscow", Birthday: "27 декабря 2000 г.", Rating: 5.0, Gender: "мужской" },
      ];
      setEntrepreneurs(staticEntrepreneurs);
    };

    fetchEntrepreneurs();
  }, []);

  return (
    <ListContainer>
      <div style={{ position: 'relative' }}>
        <ListComponent
          data={entrepreneurs}
          renderItem={(entrepreneur) => (
            <EntrepreneurItem key={entrepreneur.ID} entrepreneur={entrepreneur} />
          )}
          itemsPerPage={3}
        />
      </div>
    </ListContainer>
  );
};

export default EntrepreneursList;