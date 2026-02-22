'use client'

import { Box, Tabs } from "@chakra-ui/react";
import { ArticleFeed } from "./_components/ArticleList";

const Home = () => {
  const articles = [
    { id: "1", authorAvatar: "https://i.pravatar.cc/", authorName: "Satu Dua", title: "Ini Judul Artikel Pertama", description: "Deskripsi artikel pertama adalah sebagai berikut", thumbnail: "https://placehold.co/300x200", altThumbnail: "Ini gambar", date: "Feb 29", commentCount: "69" },
    { id: "2", authorAvatar: "https://i.pravatar.cc/", authorName: "Satu Dua", title: "Ini Judul Artikel Pertama", description: "Deskripsi artikel pertama adalah sebagai berikut", thumbnail: "https://placehold.co/300x200", altThumbnail: "Ini gambar", date: "Feb 29", commentCount: "69" },
    { id: "3", authorAvatar: "https://i.pravatar.cc/", authorName: "Satu Dua", title: "Ini Judul Artikel Pertama", description: "Deskripsi artikel pertama adalah sebagai berikut", thumbnail: "https://placehold.co/300x200", altThumbnail: "Ini gambar", date: "Feb 29", commentCount: "69" },
    { id: "4", authorAvatar: "https://i.pravatar.cc/", authorName: "Satu Dua", title: "Ini Judul Artikel Pertama", description: "Deskripsi artikel pertama adalah sebagai berikut", thumbnail: "https://placehold.co/300x200", altThumbnail: "Ini gambar", date: "Feb 29", commentCount: "69" },
    { id: "5", authorAvatar: "https://i.pravatar.cc/", authorName: "Satu Dua", title: "Ini Judul Artikel Pertama", description: "Deskripsi artikel pertama adalah sebagai berikut", thumbnail: "https://placehold.co/300x200", altThumbnail: "Ini gambar", date: "Feb 29", commentCount: "69" },
  ]

  return (
    <Box p={"4"}>
      <Tabs.Root defaultValue="first" >
        <Tabs.List>
          <Tabs.Trigger value="first">For You</Tabs.Trigger>
          <Tabs.Trigger value="second">Featured</Tabs.Trigger>
        </Tabs.List>
        <Tabs.Content value="first" overflowY={"auto"}> <ArticleFeed articles={articles} /> </Tabs.Content>
        <Tabs.Content value="second">Feed Featured Khusus</Tabs.Content>
      </Tabs.Root>
    </Box>
  );
}

export default Home;
