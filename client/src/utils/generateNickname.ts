const adjectives = [
  "친절한",
  "센스있는",
  "재미있는",
  "지적인",
  "성실한",
  "잘생긴",
  "예쁜",
  "귀여운",
  "세련된",
  "멋진",
  "진지한",
  "열정적인",
  "용감한",
  "도전적인",
  "자유로운",
  "단단한",
  "건강한",
  "명랑한",
  "매력적인",
  "카리스마 있는",
];
const animals = [
  "강아지",
  "고양이",
  "사자",
  "토끼",
  "사슴",
  "치타",
  "기린",
  "돌고래",
  "고래",
  "원숭이",
  "여우",
  "늑대",
  "코끼리",
  "코뿔소",
  "낙타",
  "팬더",
  "악어",
  "하마",
  "원앙",
  "바다표범",
];

export function getNickname() {
  const nicknameFromLocalStorage = localStorage.getItem("nickname");
  if (nicknameFromLocalStorage) {
    return nicknameFromLocalStorage;
  }
  const nickname = generateRandomNickname();
  localStorage.setItem("nickname", nickname);
  return nickname;
}

export default function generateRandomNickname() {
  const randomAdjective =
    adjectives[Math.floor(Math.random() * adjectives.length)];
  const randomAnimal = animals[Math.floor(Math.random() * animals.length)];
  return `${randomAdjective} ${randomAnimal}`;
}
