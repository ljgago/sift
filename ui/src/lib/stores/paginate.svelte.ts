class PaginateStore {
  query: string = $state("");
  page: number = $state(1);
  size: number = $state(20);
}

export const paginate = new PaginateStore();
