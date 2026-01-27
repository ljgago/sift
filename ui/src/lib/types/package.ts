export interface ResponseList {
  status: string;
  data: Pagination;
}

export interface ResponsePackage {
  status: string;
  data: string;
}

export interface Pagination {
  page: number;
  size: number;
  total_pages: number;
  total_items: number;
  items: PackageMetadata[];
}

export interface PackageMetadata {
  downloads: {
    monthly: number;
    weekly: number;
  };
  dependents: string;
  updated: string;
  searchScore: number;
  package: {
    name: string;
    keywords: string[];
    version: string;
    description: string;
    sanitized_name: string;
    publisher: {
      email: string;
      username: string;
      actor: {
        name: string;
        type: string;
        email: string;
      };
      trustedPublisher: {
        oidcConfigId: string;
        id: string;
      };
    };
    maintainers: Array<{
      email: string;
      username: string;
    }>;
    license: string;
    date: string;
    links: {
      homepage: string;
      repository: string;
      bugs: string;
      npm: string;
    };
  };
  score: {
    final: number;
    detail: {
      popularity: number;
      quality: number;
      maintenance: number;
    };
  };
  flags: {
    insecure: number;
  };
}
