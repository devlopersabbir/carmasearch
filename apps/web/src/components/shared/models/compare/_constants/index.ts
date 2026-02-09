export const exteriorColors = [
  "Beige",
  "Blue",
  "Brown",
  "Bronze",
  "Yellow",
  "Gray",
  "Green",
  "Red",
  "Black",
  "Silver",
  "Purple",
  "White",
  "Orange",
  "Gold",
];
export const interiorColors = [
  "Beige",
  "Black",
  "Gray",
  "Brown",
  "Other",
  "Blue",
  "Red",
  "Green",
  "Yellow",
  "Orange",
  "White",
];
export const interiorMaterials = [
  "Alcantara",
  "Full Leather",
  "Partial Leather",
  "Fabric",
  "Other",
  "Velour",
];
export const registrationYears = [
  "2013",
  "2014",
  "2015",
  "2016",
  "2017",
  "2018",
  "2019",
  "2020",
  "2021",
  "2022",
  "2023",
  "2024",
];

// Color mapping for swatches
export const getColorSwatch = (color: string) => {
  const colorMap: { [key: string]: string } = {
    Beige: "#F5F5DC",
    Blue: "#0000FF",
    Brown: "#8B4513",
    Bronze: "#CD7F32",
    Yellow: "#FFFF00",
    Gray: "#808080",
    Green: "#008000",
    Red: "#FF0000",
    Black: "#000000",
    Silver: "#C0C0C0",
    Purple: "#800080",
    White: "#FFFFFF",
    Orange: "#FFA500",
    Gold: "#FFD700",
    Other: "#E0E0E0",
  };
  return colorMap[color] || "#E0E0E0";
};
