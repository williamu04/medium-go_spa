'use client'

import { ArticleCard } from "@/components/ui/ArticleCard"
import { Flex, For } from "@chakra-ui/react"

type Article = {
  id: string
  authorAvatar: string
  authorName: string
  title: string
  description: string
  thumbnail: string
  altThumbnail: string
  date: string
  commentCount: string

}

interface Props {
  articles: Article[]
}

export const ArticleFeed = ({
  articles
}: Props) => {
  return (
    <Flex direction={"column"} gap={"2"}>
      <For each={articles}>
        {(article, index) => (
          <ArticleCard key={index} {...article} />
        )}
      </For>
    </Flex>
  );
}
