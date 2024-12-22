export interface Company {
  ID: string;
  OwnerID: string;
  Name: string;
  City: string;
  Description: string;
  ActivityFieldID: string;
}

export interface CompanySmall {
  ID: string;
  Name: string;
  City: string;
  ActivityFieldID: string;
}