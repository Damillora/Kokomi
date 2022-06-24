import axios from "axios";

const getSearchItems = (term) => {
    return [
        {
            url: `https://www.google.com/search?q=${term}`,
            type: `Google`,
            text: `${term}`,
            icon: `search`,
        }
    ]
}
export const getAutocompleteItems = async (term) => {
    const searchItems = getSearchItems(term);
    const allItems = [...searchItems];

    return allItems
}