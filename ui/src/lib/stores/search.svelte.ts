class SearchStore {
  name: string = $state("");

  print() {
    console.log(this.name);
  }
}

export const search = new SearchStore();
