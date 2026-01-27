export interface ResponseStars {
  status: string;
  data: Stars;
}

export interface Stars {
  repository_name: string;
  stars_count: number;
  updated_at: string;
}
