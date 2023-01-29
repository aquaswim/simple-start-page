export interface ILink {
  icon?: string;
  url?: string;
  name: string;
}

export interface ISetting {
  siteTitle: string;
  siteAbout: string;
  hideSetting: boolean;
}

export interface ILoginResp {
  token: string;
}

export interface IErrorState {
  message: string;
}

export interface IResponse<T> {
  data?: T;
  code: number;
  message: string;
}
