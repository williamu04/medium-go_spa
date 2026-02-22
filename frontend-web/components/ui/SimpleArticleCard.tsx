import { Avatar, Card, Heading, HStack, Text } from "@chakra-ui/react"


interface Props {
  authorAvatar: string
  authorName: string
  title: string
  date: string
}

export const SimpleArticleCard = ({
  authorAvatar, authorName, title, date
}: Props) => {
  return (
    <Card.Root variant={"outline"}>
      <Card.Header>
        <HStack>
          <Avatar.Root>
            <Avatar.Image src={authorAvatar} />
            <Avatar.Fallback name={authorName} />
          </Avatar.Root>
          <Text>{authorName}</Text>
        </HStack>
      </Card.Header>
      <Card.Body>
        <Heading>{title}</Heading>
      </Card.Body>
      <Card.Footer>
        <Text>{date}</Text>
      </Card.Footer>
    </Card.Root>
  )
}
